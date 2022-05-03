import express from 'express';
import userRoutes from './users';
import hello from './hello';

const router = express.Router();

router.use('/user', userRoutes);
router.use('/hello', hello);

export default router;
